import * as fs from 'fs';
import { z } from 'zod';

interface Config {
    BACKEND_URL: string;
    ADMIN_USER: string;
    ADMIN_PASSWORD: string;
}

export let config: Config = {} as never;

const loadConfig = () => {
    const configFile = process.env.CONFIG_FILE;
    if (configFile == undefined) {
        throw new Error("CONFIG_FILE is not set");
    }

    const configContents = fs.readFileSync(configFile, 'utf8');
    const schema = z.object({
        BACKEND_URL: z.string(),
        ADMIN_USER: z.string(),
        ADMIN_PASSWORD: z.string(),
    }).strict();
    config = schema.parse(JSON.parse(configContents));
};

loadConfig();