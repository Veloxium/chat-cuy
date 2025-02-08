import { z } from 'zod';

const registerSchema = z.object({
    username: z.string().min(2, 'Username minimal 2 karakter'),
    email: z.string().min(6, 'Email minimal 6 karakter').email('Email tidak valid'),
    password: z.string().min(6, 'Password minimal 6 karakter'),
    konfirmasiPassword: z.string().min(6, 'Konfirmasi Password minimal 6 karakter'),
}).refine(data => data.password === data.konfirmasiPassword, {
    message: 'Password dan Konfirmasi Password harus sama',
    path: ['konfirmasiPassword'],
});

export default registerSchema;
