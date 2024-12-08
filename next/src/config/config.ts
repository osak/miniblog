import { z } from 'zod';

interface Config {
    BACKEND_URL: string;
    ADMIN_USER: string;
    ADMIN_PASSWORD: string;
}

export let config: Config = {} as never;

const loadConfig = () => {
    const schema = z.object({
        BACKEND_URL: z.string(),
        ADMIN_USER: z.string(),
        ADMIN_PASSWORD: z.string(),
    });
    config = schema.parse(process.env);
};

loadConfig();