import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { BankAccount } from 'src/models/bank-account.model';
import { Repository } from 'typeorm';

@Injectable()
export class BankAccountService {
  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepository: Repository<BankAccount>
  ) {}

  findAll(): Promise<BankAccount[]> {
    return this.bankAccountRepository.find();
  }

  findOne(id: string): Promise<BankAccount> {
    return this.bankAccountRepository.findOneOrFail(id);
  }
}
