import fs from 'fs/promises';
import path from 'path';
import type { Payload } from './prometheus/client';
import { Mutex } from 'async-mutex';

const FILE_PATH = path.join('./files', 'saved.txt');
const DIR_PATH = './files';

const mutex = new Mutex();

export async function updateNumber(payload: Payload): Promise<void> {
    const release = await mutex.acquire();
    try {
        await ensureDirectoryExists(DIR_PATH);
        
        let currentValue = 0;
        try {
            const data = await fs.readFile(FILE_PATH, 'utf8');
            currentValue = parseInt(data.trim(), 10);
            
            if (isNaN(currentValue)) {
                console.warn('File contained invalid number, resetting to 0');
                currentValue = 0;
            }
            
            console.log('Current value:', currentValue);
        } catch (error) {
            console.log('No existing file found, creating new one');
        }
        
        const newValue = currentValue + payload.value;
        console.log('Updated value:', newValue);
        
        await fs.writeFile(FILE_PATH, newValue.toString());
        console.log('File saved successfully');
    } catch (error) {
        console.error('Error updating number:', error);
        throw error; 
    }finally{
        release(); 
    }
}


async function ensureDirectoryExists(dirPath: string): Promise<void> {
    try {
        await fs.mkdir(dirPath, { recursive: true });
    } catch (error: any) {
        if (error.code !== 'EEXIST') {
            throw error;
        }
    }
}


export async function readFile(): Promise<number> {    
    try {
        const data = await fs.readFile(FILE_PATH, 'utf8');
        return parseInt(data.trim(), 10);
    } catch (error) {
        console.error('Error reading file:', error);
        throw error;
    }
    
}