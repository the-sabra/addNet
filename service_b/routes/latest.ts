import express from 'express';
import { readFile } from '../save-file';

export const latestRoute = express.Router();

latestRoute.get('/latest', async (req, res) => {
  try {
    const latestValue = await readFile();
    console.log('Latest value:', latestValue); // Log the latest value to the console
    res.status(200).json({ latestValue });
  } catch (error) {
    console.error('Error reading file:', error);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});