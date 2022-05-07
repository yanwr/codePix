import { Controller, Get, Param, ParseUUIDPipe, Post } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { PixKey } from 'src/models/pix-key.model';
import { Repository } from 'typeorm';

@Controller('pix-keys/:bankAccountId')
export class PixKeyController {
  constructor(
    @InjectRepository(PixKey)
    private pixKeyRepository: Repository<PixKey>
  ) {}

  @Get()
  getAllPixKeys(@Param("bankAccountId", new ParseUUIDPipe({ version: "4"})) bankAccountId: string) {
    return this.pixKeyRepository.find({
      where: { bankAccountId },
      order: { createdAt: "DESC" }
    });
  }

  @Post()
  createAndSavaPixKey() {

  }

  @Get("exists")
  isTherePixKey() {

  }

}
