import { Body, Controller, Get, Inject, InternalServerErrorException, Param, ParseUUIDPipe, Post, UnprocessableEntityException, ValidationPipe } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { InjectRepository } from '@nestjs/typeorm';
import PixKeyDTO from 'src/dto/pix-key.dto';
import { PixService } from 'src/grpc/types/pix-service.grpc.types';
import { BankAccount } from 'src/models/bank-account.model';
import { PixKey } from 'src/models/pix-key.model';
import { Repository } from 'typeorm';

@Controller('pix-keys/:bankAccountId')
export class PixKeyController {
  constructor(
    @InjectRepository(PixKey)
    private pixKeyRepository: Repository<PixKey>,
    @InjectRepository(PixKey)
    private bankAccountRepository: Repository<BankAccount>,
    @Inject("CODEPIX_PACKAGE")
    private clientGRPC: ClientGrpc
  ) {}

  @Get()
  getAllPixKeys(@Param("bankAccountId", new ParseUUIDPipe({ version: "4"})) bankAccountId: string) {
    return this.pixKeyRepository.find({
      where: { bankAccountId },
      order: { createdAt: "DESC" }
    });
  }

  @Post()
  async createAndSavePixKey(
    @Param("bankAccountId", new ParseUUIDPipe({ version: "4" })) bankAccountId: string, 
    @Body(new ValidationPipe({ errorHttpStatusCode: 422 })) pixKeyBody: PixKeyDTO
  ) {
    await this.bankAccountRepository.findOneOrFail(bankAccountId);
    const pixKeyService: PixService = this.clientGRPC.getService("PixServiceController");
    const pixKeyNotFound = await this.checkPixKeyNotFound(pixKeyBody, pixKeyService);
    if(!pixKeyNotFound) {
      throw new UnprocessableEntityException("PixKey already exists");
    }
    const createdPixKey = await pixKeyService.registerPixKey({ ...pixKeyBody, accountId: bankAccountId}).toPromise();
    if(createdPixKey.error) {
      throw new InternalServerErrorException(createdPixKey.error);
    }
    const pixKey = this.pixKeyRepository.create({
      id: createdPixKey.id,
      bankAccountId: bankAccountId,
      ...pixKeyBody
    });
    return await this.pixKeyRepository.save(pixKey);
  }

  async checkPixKeyNotFound(param: { key: string, kind: string}, pixKeyService: PixService): Promise<boolean> {
    try {
      await pixKeyService.find(param).toPromise();
      return false;
    } catch (e) {
      if (e.details === "no key was found") {
        return true;
      }

      throw new InternalServerErrorException("Server not available");
    }
  }

  @Get("exists")
  isTherePixKey() {

  }

}
