import { Injectable, InternalServerErrorException, UnprocessableEntityException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import PixKeyDTO from 'src/dto/pix-key.dto';
import { GrpcServices } from 'src/grpc/grpc.service';
import { PixService } from 'src/grpc/types/pix-service.grpc.types';
import { PixKey } from 'src/models/pix-key.model';
import { Repository } from 'typeorm';
import { BankAccountService } from '../bank-account/bank-account.service';

@Injectable()
export class PixKeyService {
  constructor(
    @InjectRepository(PixKey)
    private pixKeyRepository: Repository<PixKey>,
    private bankAccountService: BankAccountService,
    private grpcService: GrpcServices
  ) {}

  findAll(bankAccountId: string): Promise<PixKey[]> {
    return this.pixKeyRepository.find({
      where: { bankAccountId },
      order: { createdAt: "DESC" }
    });
  }

  async createAndSave(bankAccountId: string, pixKeyDTO: PixKeyDTO): Promise<PixKey> {
    return await this.save(await this.create(bankAccountId, pixKeyDTO));
  }

  async create(bankAccountId: string, pixKeyDTO: PixKeyDTO): Promise<PixKey> {
    await this.bankAccountService.findOneOrThrow(bankAccountId);
    const grpcPixKeyService: PixService = this.grpcService.getPixService();
    
    if(!(await this.isValidPixKey(pixKeyDTO, grpcPixKeyService))) {
      throw new UnprocessableEntityException("PixKey already exists");
    }

    const createdPixKeyByCodePixService = await grpcPixKeyService.registerPixKey({ ...pixKeyDTO, accountId: bankAccountId}).toPromise();
    if(createdPixKeyByCodePixService.error) {
      throw new InternalServerErrorException(createdPixKeyByCodePixService.error);
    }

    return this.pixKeyRepository.create({
      id: createdPixKeyByCodePixService.id,
      bankAccountId: bankAccountId,
      ...pixKeyDTO
    });
  }

  async save(pixKey: PixKey): Promise<PixKey> {
    return await this.pixKeyRepository.save(pixKey);
  }

  async isValidPixKey(param: { key: string, kind: string}, grpcPixKeyService: PixService): Promise<boolean> {
    try {
      await grpcPixKeyService.find(param).toPromise();
      return true;
    } catch (e) {
      if (e.details === "no key was found") {
        return false;
      }
      throw new InternalServerErrorException("Server not available");
    }
  }
}
