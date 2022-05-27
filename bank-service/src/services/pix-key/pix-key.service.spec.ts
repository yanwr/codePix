import { Test, TestingModule } from '@nestjs/testing';
import { PixKeyService } from './pix-key.service';

describe('PixKeyService', () => {
  let service: PixKeyService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [PixKeyService],
    }).compile();

    service = module.get<PixKeyService>(PixKeyService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
