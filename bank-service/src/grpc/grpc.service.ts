import { Inject, Injectable } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { PixService } from './types/pix-service.grpc.types';

@Injectable()
export class GrpcServices {

  private PIX_SERVICE: string = "PixServiceController";

  constructor(
    @Inject("CODEPIX_PACKAGE")
    private clientGRPC: ClientGrpc
  ) {}
  
  getPixService(): PixService {
    return this.clientGRPC.getService(this.PIX_SERVICE);
  }
}
