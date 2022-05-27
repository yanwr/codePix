import { NestFactory } from '@nestjs/core';
import { Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';
import { ModelNotFoundExceptionFilter } from './exceptions/model-not-found.exception';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix("api");

  app.useGlobalFilters(new ModelNotFoundExceptionFilter());
  app.connectMicroservice({
    transport: Transport.KAFKA,
    options: {
      client: {
        brokers: [`${process.env.KAFKA_BROKER}:${process.env.KAFKA_BROKER_PORT}`]
      },
      consumer: {
        groupId: !process.env.KAFKA_CONSUMER_GROUP_ID || process.env.KAFKA_CONSUMER_GROUP_ID === "" ? "my-consumer-" + Math.random() : process.env.KAFKA_CONSUMER_GROUP_ID
      }
    }
  });
  app.startAllMicroservices();
  await app.listen(3000);
}
bootstrap();