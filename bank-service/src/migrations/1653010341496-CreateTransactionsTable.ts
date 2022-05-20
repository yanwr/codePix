import {MigrationInterface, QueryRunner, Table, TableForeignKey} from "typeorm";

export class CreateTransactionsTable1612394636514 implements MigrationInterface {
  
    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
          new Table({
            name: 'transactions',
            columns: [
              {
                name: 'id',
                type: 'uuid',
                isPrimary: true,
              },
              {
                name: 'externalId',
                type: 'uuid',
              },
              {
                name: 'amount',
                type: 'double precision',
              },
              {
                name: 'description',
                type: 'varchar',
                isNullable: true
              },
              {
                name: 'bankAccountId',
                type: 'uuid',
              },
              {
                name: 'bankAccountFromId',
                type: 'uuid',
                isNullable: true
              },
              {
                name: 'pixKey',
                type: 'varchar',
                isNullable: true
              },
              {
                name: 'pixKeyKind',
                type: 'varchar',
                isNullable: true
              },
              {
                name: 'status',
                type: 'varchar',
              },
              {
                name: 'operation',
                type: 'varchar',
              },
              {
                name: 'createdAt',
                type: 'timestamp',
                default: 'CURRENT_TIMESTAMP',
              },
            ],
          }),
        );
    
        await queryRunner.createForeignKey(
          'transactions',
          new TableForeignKey({
            name: 'transactionsBankAccountIdForeignKey',
            columnNames: ['bankAccountId'],
            referencedColumnNames: ['id'],
            referencedTableName: 'bankAccounts',
          }),
        );
    
      }
    
      public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.dropForeignKey(
          'transactions',
          'transactionsBankAccountIdForeignKey',
        );
        await queryRunner.dropTable('transactions');
      }
}