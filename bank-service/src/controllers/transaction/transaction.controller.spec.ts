import { Test, TestingModule } from '@nestjs/testing';
import { TransactionControllerController } from './transaction.controller';

describe('TransactionControllerController', () => {
  let controller: TransactionControllerController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [TransactionControllerController],
    }).compile();

    controller = module.get<TransactionControllerController>(TransactionControllerController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
