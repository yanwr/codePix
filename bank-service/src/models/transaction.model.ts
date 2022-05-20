import {
  BeforeInsert,
  Column,
  CreateDateColumn,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { BankAccount } from './bank-account.model';
import { PixKeyKind } from './pix-key.model';
import { v4 as uuidv4 } from 'uuid';

export enum TransactionStatus{
  TRANSACTION_PENDING   = "PENDING",
	TRANSACTION_COMPLETED = "COMPLETED",
	TRANSACTION_CANCELED  = "CANCELED",
	TRANSACTION_CONFIRMED = "CONFIRMED"
};

export enum TransactionOperation {
  DEBIT = 'DEBIT',
  CREDIT = 'CREDIT'
}

@Entity({ name: 'transactions' })
export class Transaction {
  @PrimaryGeneratedColumn("uuid")
  id: string;

  @Column()
  externalId: string;

  @Column()
  amount: number;

  @Column()
  description: string;

  @ManyToOne(() => BankAccount)
  @JoinColumn({name: "bankAccountId"})
  bankAccount: BankAccount;

  @Column()
  bankAccountId: string;

  @ManyToOne(() => BankAccount)
  @JoinColumn({name: "bankAccountFromId"})
  bankAccountFrom: BankAccount;

  @Column()
  bankAccountFromId: string;

  @Column()
  pixKey: string;

  @Column()
  pixKeyKind: PixKeyKind;

  @Column()
  status: TransactionStatus = TransactionStatus.TRANSACTION_PENDING;

  @Column()
  operation: TransactionOperation

  @CreateDateColumn({ type: 'timestamp' })
  createdAt: Date;

  @BeforeInsert() generateId() {
    if (this.id) {
      return;
    }
    this.id = uuidv4();
  }

  @BeforeInsert() generateExternalId() {
    if (this.externalId) {
      return;
    }
    this.externalId = uuidv4();
  }
}