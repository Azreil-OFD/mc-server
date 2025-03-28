import sqlite3 from 'sqlite3';
import { open, Database } from 'sqlite';

export class sqliteDb {
  private dbPath: string;
  private db: Database | null;

  constructor(dbPath: string) {
    this.dbPath = dbPath;
    this.db = null;
  }

  private async connect(): Promise<void> {
    if (!this.db) {
      this.db = await open({
        filename: this.dbPath,
        driver: sqlite3.Database
      });
    }
  }

  async getPasswordByUsername(username: string): Promise<string | null> {
    await this.connect();

    try {
      const result = await this.db!.get<{ password: string }>(
        'SELECT password FROM authme WHERE username = ?',
        [username]
      );

      return result ? result.password : null;
    } catch (err) {
      console.error('Ошибка при выполнении запроса:', err);
      return null;
    }
  }

  async close(): Promise<void> {
    if (this.db) {
      await this.db.close();
      this.db = null;
    }
  }
}