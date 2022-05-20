import { IsIn, IsNotEmpty, IsNumber, IsOptional, IsString, Min } from "class-validator";
import { PixKeyKind } from "src/models/pix-key.model";

export default class TransactionDTO {
  @IsString()
  @IsNotEmpty()
  pixKey: string;

  @IsString()
  @IsIn(Object.values(PixKeyKind))
  @IsNotEmpty()
  pixKind: PixKeyKind;

  @IsString()
  @IsOptional()
  description: string = null;

  @IsNumber()
  @Min(0.01)
  @IsNotEmpty()
  readonly amount: number;
}