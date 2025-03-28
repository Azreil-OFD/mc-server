import type { PlayerData } from "~/types";
import { getOps } from "./users/getOps"


export class JsDB {
    filepath: string = '../ops.json'

    async opsByUsername(username: string): Promise<boolean | null> {
        let player = await getOps(this.filepath)
        
        const findPlayer = player.find(e => {
            if (e.name === username)
                return true
            return false
        })
        
        if(findPlayer === undefined) 
            return false
        return true
    }
}