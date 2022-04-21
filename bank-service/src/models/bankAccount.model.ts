import {BeforeInsert, Column, CreateDateColumn, Entity, PrimaryGeneratedColumn} from "typeorm";
import { v4 as uuidV4 } from "uuid"

@Entity({ name: "bankAccounts" })
export class BankAccount {
  @PrimaryGeneratedColumn("uuid")
  id: string;

  @Column()
  accountNumber: string;

  @Column()
  ownerName: string;

  @Column()
  balance: number;

  @CreateDateColumn({ type: "timestamp" })
  createdAt: Date;

  @BeforeInsert()
  generateId() {
    if(this.id) {
      return;
    }
    this.id = uuidV4();
  }

  @BeforeInsert()
  generateBalance() {
    if(this.balance) {
      return;
    }
    this.balance = 0;
  }
}
