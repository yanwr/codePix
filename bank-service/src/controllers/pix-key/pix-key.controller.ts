import { 
  Body,
  Controller, 
  Get, 
  Param, 
  ParseUUIDPipe, 
  Post, 
  ValidationPipe 
} from '@nestjs/common';
import PixKeyDTO from 'src/dto/pix-key.dto';
import { PixKeyService } from 'src/services/pix-key/pix-key.service';

@Controller('pix-keys/:bankAccountId')
export class PixKeyController {
  constructor(
    private pixKeyService: PixKeyService
  ) {}

  @Get()
  getAllPixKeys(@Param("bankAccountId", new ParseUUIDPipe({ version: "4"})) bankAccountId: string) {
    return this.pixKeyService.findAll(bankAccountId);
  }

  @Post()
  async createAndSavePixKey(
    @Param("bankAccountId", new ParseUUIDPipe({ version: "4" })) bankAccountId: string, 
    @Body(new ValidationPipe({ errorHttpStatusCode: 422 })) pixKeyDTO: PixKeyDTO
  ) {
    return await this.pixKeyService.createAndSave(bankAccountId, pixKeyDTO);
  }

}
