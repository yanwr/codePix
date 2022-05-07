import {BeforeInsert, Column, CreateDateColumn, Entity, JoinColumn, ManyToOne, PrimaryGeneratedColumn} from "typeorm";
import { BankAccount } from "./bank-account.model";
import { v4 as uuidV4 } from "uuid";

export enum PixKeyKind {
  CPF = "CPF",
  EMAIL = "email"
}

@Entity({ name: "pixKeys" })
export class PixKey {
  @PrimaryGeneratedColumn("uuid")
  id: string;

  @Column()
  kind: PixKeyKind;

  @Column()
  key: string;

  @ManyToOne(() => BankAccount)
  @JoinColumn({ name: "bankAccountId" })
  bankAccount: BankAccount;

  @Column()
  bankAccountId: string;

  @CreateDateColumn({ type: "timestamp" })
  createdAt: Date;

  @BeforeInsert()
  generateId() {
    if(this.id) {
      return;
    }
    this.id = uuidV4();
  }
}
