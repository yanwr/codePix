import {MigrationInterface, QueryRunner, Table, TableForeignKey} from "typeorm";

export class CreatePixKeyTable1651947499799 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
            new Table({
                name: "pixKeys",
                columns: [
                    {
                        name: "id",
                        type: "uuid",
                        isPrimary: true
                    },
                    {
                        name: "kind",
                        type: "varchar"
                    },
                    {
                        name: "key",
                        type: "varchar"
                    },
                    {
                        name: "bankAccountId",
                        type: "uuid"
                    },
                    {
                        name: "createdAt",
                        type: "timestamp",
                        default: "CURRENT_TIMESTAMP"
                    }
                ]
            })
        );

        await queryRunner.createForeignKey("pixKeys", new TableForeignKey({
            name: "pixKeysBankAccountIdForeignKey",
            columnNames: ["bankAccountId"],
            referencedColumnNames: ["id"],
            referencedTableName: "bankAccounts"
        }));
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.dropForeignKey(
            "pixKeys",
            "pixKeysBankAccountIdForeignKey"
        );
        await queryRunner.dropTable("pixKeys");
    }

}
