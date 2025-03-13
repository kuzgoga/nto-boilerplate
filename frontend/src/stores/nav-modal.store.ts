import { defineStore } from "pinia";

export interface NavModalState {
    visible: boolean
}

export const useNavModalStore = defineStore('nav-modal', {
    state: (): NavModalState => ({
        visible: false
    }),
    actions: {
        changeVisibility() {
            this.visible = !this.visible
        }
    }
})