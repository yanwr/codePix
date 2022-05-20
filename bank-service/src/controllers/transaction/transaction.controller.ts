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
export class TransactionController {

  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepository: Repository<BankAccount>,
    @InjectRepository(Transaction)
    private transactionRepository: Repository<Transaction>,
    @InjectRepository(PixKey)
    private pixKeyRepository: Repository<PixKey>,
  ) {}

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
    @Param(
      'bankAccountId',
      new ParseUUIDPipe({ version: '4', errorHttpStatusCode: 422 }),
    )
    bankAccountId: string,
    @Body(new ValidationPipe({ errorHttpStatusCode: 422 }))
    transactionBody: TransactionDTO,
  ) {
    await this.bankAccountRepository.findOneOrFail(bankAccountId);
  }
}