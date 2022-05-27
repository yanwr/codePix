import { Controller, Get, Param, ParseUUIDPipe } from '@nestjs/common';
import { BankAccountService } from 'src/services/bank-account/bank-account.service';

@Controller('bank-accounts')
export class BankAccountController {
  constructor(
    private bankAccountService: BankAccountService
  ) {}

  @Get()
  getAllBankAccounts() {
    return this.bankAccountService.findAll();
  }
  
  @Get(":bankAccountId")
  getBankAccount(@Param("bankAccountId", new ParseUUIDPipe({ version: "4"})) bankAccountId: string) {
    return this.bankAccountService.findOneOrThrow(bankAccountId);
  }
}
