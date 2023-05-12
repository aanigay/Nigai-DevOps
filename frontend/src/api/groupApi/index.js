import {GroupAPIInstance} from "@/api";

export const GroupApi = {
    getRoleById(id) {
        const url = '/getRoleById'
        return GroupAPIInstance.get(url, {
            params: {
                id: id
            }
        })
    },
    changeRoleById(id, role) {
        const url = '/changeRole'
        const data = {id, role}
        return GroupAPIInstance.put(url, data)
    }
}