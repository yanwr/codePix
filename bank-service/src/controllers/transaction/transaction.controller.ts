import {
  Body,
  Controller,
  Get,
  Inject,
  OnModuleDestroy,
  OnModuleInit,
  Param,
  ParseUUIDPipe,
  Post,
  ValidationPipe,
} from '@nestjs/common';
import { ClientKafka, MessagePattern, Payload } from '@nestjs/microservices';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectRepository } from '@nestjs/typeorm';
import TransactionDTO from 'src/dto/transaction';
import { BankAccount } from 'src/models/bank-account.model';
import { PixKey } from 'src/models/pix-key.model';
import {
  Transaction,
  TransactionOperation,
  TransactionStatus,
} from 'src/models/transaction.model';
import { Repository } from 'typeorm';

@Controller('transactions/:bankAccountId')
export class TransactionController implements OnModuleInit, OnModuleDestroy {

  private kafkaProducer: Producer;

  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepository: Repository<BankAccount>,
    @InjectRepository(Transaction)
    private transactionRepository: Repository<Transaction>,
    @InjectRepository(PixKey)
    private pixKeyRepository: Repository<PixKey>,
    @Inject("TRANSACTION_SERVICE")
    private kafkaClient: ClientKafka
  ) {}

  async onModuleInit() {
    this.kafkaProducer = await this.kafkaClient.connect();
  }

  async onModuleDestroy() {
    await this.kafkaProducer.disconnect();
  }

  @Get()
  getAllTransactions(
    @Param('bankAccountId', new ParseUUIDPipe({ version: '4', errorHttpStatusCode: 422 })) bankAccountId: string,
  ) {
    return this.transactionRepository.find({
      where: { bankAccountId: bankAccountId },
      order: { createdAt: 'DESC' }
    });
  }

  @Post()
  async createAndSaveTransaction(
    @Param('bankAccountId', new ParseUUIDPipe({ version: '4', errorHttpStatusCode: 422 }))
    bankAccountId: string,
    @Body(new ValidationPipe({ errorHttpStatusCode: 422 }))
    transactionBody: TransactionDTO,
  ) {
    await this.bankAccountRepository.findOneOrFail(bankAccountId);

    let transaction = this.transactionRepository.create({
      ...transactionBody,
      amount: transactionBody.amount * -1,
      bankAccountId: bankAccountId,
      operation: TransactionOperation.DEBIT
    });

    transaction = await this.bankAccountRepository.save(transaction);

    const sendData = {
      id: transaction.externalId,
      accountId: bankAccountId,
      amount: transaction.amount,
      pixKeyTo: transaction.pixKey,
      pixKeyKind: transaction.pixKeyKind,
      description: transaction.description
    };

    await this.kafkaProducer.send({
      topic: "transactions",
      messages: [
        { key: "transactions", value: JSON.stringify(sendData) }
      ]
    });

    return transaction;
  }

  @MessagePattern("bank001")
  async doTransaction(@Payload() data) {
    switch (data.value.status) {
      case TransactionStatus.TRANSACTION_PENDING:
        await this.receivedTransaction(data.value);
        break;
      case TransactionStatus.TRANSACTION_CONFIRMED:
        await this.confimedTransaction(data.value);
        break;
      default:
        break;
    }
  }

  async receivedTransaction(data) {
    const pixKey = await this.pixKeyRepository.findOneOrFail({
      where: {
        key: data.pixKeyTo,
        kind: data.pixKeyKindTo
      }
    });

    const transaction = this.transactionRepository.create({
      externalId: data.id,
      amount: data.amount,
      description: data.description,
      bankAccountId: data.bankAccountId,
      bankAccountFromId: data.accountId,
      operation: TransactionOperation.CREDIT,
      status: TransactionStatus.TRANSACTION_COMPLETED
    });

    this.transactionRepository.save(transaction);

    const sendData = {
      ...data,
      status: TransactionStatus.TRANSACTION_CONFIRMED
    };

    await this.kafkaProducer.send({
      topic: "transaction_confirmation",
      messages: [
        { key: "transaction_confirmation", value: JSON.stringify(sendData) }
      ]
    });
  }

  async confimedTransaction(data) {
    const transaction = await this.transactionRepository.findOneOrFail({
      where: {
        externalId: data.id
      }
    });

    await this.transactionRepository.update(
      { id: data.id },
      { status: TransactionStatus.TRANSACTION_COMPLETED }
    );

    const sendData = {
      id: data.id,
      accountId: transaction.bankAccountId,
      amount: Math.abs(transaction.amount),
      pixKeyTo: transaction.pixKey,
      pixKeyKind: transaction.pixKeyKind,
      description: transaction.description,
      status: TransactionStatus.TRANSACTION_COMPLETED
    };

    await this.kafkaProducer.send({
      topic: "transaction_confirmation",
      messages: [
        { key: "transaction_confirmation", value: JSON.stringify(sendData) }
      ]
    });
  }
}