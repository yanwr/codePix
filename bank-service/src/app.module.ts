import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BankAccount } from './models/bank-account.model';
import { PixKey } from './models/pix-key.model';
import { BankAccountController } from './controllers/bank-account/bank-account.controller';
import { TransactionController } from './controllers/transaction/transaction.controller';
import { PixKeyController } from './controllers/pix-key/pix-key.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { Transaction } from './models/transaction.model';
import { BankAccountService } from './services/bank-account/bank-account.service';
import { PixKeyService } from './services/pix-key/pix-key.service';
import { TransactionService } from './services/transaction/transaction.service';
import { GrpcServices } from './grpc/grpc.service';

@Module({
  imports: [
    ConfigModule.forRoot(),
    TypeOrmModule.forRoot({
      type: process.env.TYPEORM_CONNECTION as any,
      host: process.env.TYPEORM_HOST,
      port: parseInt(process.env.TYPEORM_PORT),
      username: process.env.TYPEORM_USERNAME,
      password: process.env.TYPEORM_PASSWORD,
      database: process.env.TYPEORM_DATABASE,
      entities: [BankAccount, PixKey, Transaction]
    }),
    TypeOrmModule.forFeature([BankAccount, PixKey, Transaction]),
    ClientsModule.register([
      {
        name: "CODEPIX_PACKAGE",
        transport: Transport.GRPC,
        options: {
          url: `${process.env.GRPC_URL}:${process.env.GRPC_PORT}`,
          package: "codePix",
          protoPath: [join(__dirname, "protofiles/pixkey.proto")]
        }
      }
    ]),
    ClientsModule.register([
      {
        name: "TRANSACTION_SERVICE",
        transport: Transport.KAFKA,
        options: {
          client: {
            clientId: process.env.KAFKA_CLIENT_ID,
            brokers: [`${process.env.KAFKA_BROKER}:${process.env.KAFKA_BROKER_PORT}`]
          },
          consumer: {
            groupId: !process.env.KAFKA_CONSUMER_GROUP_ID || process.env.KAFKA_CONSUMER_GROUP_ID === '' ? "my-consumer-" + Math.random() : process.env.KAFKA_CONSUMER_GROUP_ID
          }
        }
      }
    ])
  ],
  controllers: [AppController, BankAccountController, PixKeyController, TransactionController],
  providers: [AppService, GrpcServices, BankAccountService, PixKeyService, TransactionService],
})
export class AppModule {}
