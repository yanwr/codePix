import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ModelNotFoundExceptionFilter } from './exceptions/model-not-found.exception';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix("api");

  app.useGlobalFilters(new ModelNotFoundExceptionFilter());

  await app.listen(3000);
}
bootstrap();