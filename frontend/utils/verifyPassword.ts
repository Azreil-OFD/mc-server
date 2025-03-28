import crypto from 'crypto';

/**
 * Проверка пароля на валидность (формат: $SHA$salt$hash)
 * @param {string} password - Пароль для проверки
 * @param {string} hashedPassword - Хеш в формате AuthMe SHA256
 * @returns {boolean}
 */
export function verifyPassword(password: string, hashedPassword: string): { verify: Boolean, salt: string } {
    const parts = hashedPassword.split('$');

    if (parts.length !== 4 || parts[1] !== 'SHA') {
        return {verify: false, salt: ''};
    }

    const salt = parts[2];
    const hashToVerify = parts[3];

    const passwordHash = crypto.createHash('sha256').update(password).digest('hex');
    const finalHash = crypto.createHash('sha256').update(passwordHash + salt).digest('hex');

    return {
        verify: finalHash === hashToVerify, salt
    };
}
