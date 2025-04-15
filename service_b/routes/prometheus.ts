import express from 'express';
import promClient from '../prometheus/client';

export const prometheusRoute = express.Router();

prometheusRoute.get('/metrics', async (req, res) => {
    try {
        res.set('Content-Type', promClient.register.contentType);
        res.end(await promClient.register.metrics());
    } catch (error) {
        console.error('Error fetching metrics:', error);
        res.status(500).json({ error: 'Internal Server Error' });
    }
    }
);