import { z } from 'zod';

const loginSchema = z.object({
    email: z.string().min(6, 'Email minimal 6 karakter').email('Email tidak valid'),
    password: z.string().min(6, 'Password minimal 6 karakter'),
});

export default loginSchema;

