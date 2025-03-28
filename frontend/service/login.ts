import { JsDB } from "~/repository/js";
import { sqliteDb } from "~/repository/sqlite";
import { verifyPassword } from "~/utils/verifyPassword";
import MD5 from "crypto-js/md5";
const jsonDatabase = new JsDB()
const sqliteDatabase = new sqliteDb('../plugins/AuthMe/authme.db');

export const login = async (login: string, password: string) => {
    const passwordByDatabase = await sqliteDatabase.getPasswordByUsername(login)
    if (passwordByDatabase === null) {
        return {
            status: false,
            data: {},
            error: {
                code: 401,
                message: 'Не верный логин или пароль!'
            }
        }
    }

    const { verify, salt } = verifyPassword(password, passwordByDatabase)
    if (!verify) {
        return {
            status: false,
            data: {},
            error: {
                code: 401,
                message: 'Не верный логин или пароль!'
            }
        }
    }
    const isOps = jsonDatabase.opsByUsername(login)
    return {
        status: true,
        data: {
            username: login,
            token: login + ":" + MD5(salt).toString(),
        },
        error: {}
    }

}