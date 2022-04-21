import { MigrationInterface, QueryRunner, Table } from "typeorm"

export class CreateBankAccountsTable1650481671863 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
            new Table({
                 name: "bankAccounts",
                 columns: [
                     {
                         name: "id",
                         type: "uuid",
                         isPrimary: true
                     },
                     {
                         name: "accountNumber",
                         type: "varchar"
                     },
                     {
                         name: "ownerName",
                         type: "varchar"
                     },
                     {
                         name: "balance",
                         type: "double precision"
                     },
                     {
                         name: "createdAt",
                         type: "timestamp",
                         default: "CURRENT_TIMESTAMP"
                     }
                 ]
            })
        )
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.dropTable("bankAccounts")
    }

}