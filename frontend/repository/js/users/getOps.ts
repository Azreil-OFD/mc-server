import { promises as fs } from 'fs';
import type { PlayerData } from '~/types';



// Функция для чтения данных из файла
export async function getOps(filePath: string): Promise<PlayerData[]> {
  try {
    const data = await fs.readFile(filePath, 'utf8');
    const parsedData: PlayerData[] = JSON.parse(data);

    return parsedData;
  } catch (error) {
    console.error('Ошибка при чтении файла:', error);
    return []
  }
}