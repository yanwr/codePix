import { Controller, Get, Param, ParseUUIDPipe } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { BankAccount } from 'src/models/bank-account.model';
import { Repository } from 'typeorm';

@Controller('bank-accounts')
export class BankAccountController {
  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepository: Repository<BankAccount>
  ) {}

  @Get()
  getAllBankAccounts() {
    return this.bankAccountRepository.find();
  }
  
  @Get(":bankAccountId")
  getBankAccount(@Param("bankAccountId", new ParseUUIDPipe({ version: "4"})) bankAccountId: string) {
    return this.bankAccountRepository.findOneOrFail(bankAccountId);
  }
}
