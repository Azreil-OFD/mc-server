import MD5 from "crypto-js/md5";
import { sqliteDb } from "~/repository/sqlite";

const sqliteDatabase = new sqliteDb('../plugins/AuthMe/authme.db');

export const authMiddleware = async (token: string) => {
    const [username, tokenHash] = token.split(":");
    if (!username || !tokenHash) {
        return {
            status: false,
            error: {
                code: 401,
                message: 'Неверный формат токена!'
            }
        };
    }

    const passwordByDatabase = await sqliteDatabase.getPasswordByUsername(username);

    if (!passwordByDatabase) {
        return {
            status: false,
            error: {
                code: 401,
                message: 'Пользователь не найден!'
            }
        };
    }

    const salt = passwordByDatabase.split('$')[2];
    const expectedHash = MD5(salt).toString();

    if (tokenHash !== expectedHash) {
        return {
            status: false,
            error: {
                code: 401,
                message: 'Невалидный токен авторизации!'
            }
        };
    }

    return {
        status: true,
        data: {
            username
        },
        error: {}
    };
};
