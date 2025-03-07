export type TableEmits = {
    (e: 'onCreateOpen'): void
    (e: 'onCreateClose', data: any): void
    (e: 'onUpdateOpen'): void
    (e: 'onUpdateClose', data: any): void
    (e: 'onOpen'): void
    (e: 'onClose', data: any): void
    (e: 'onDelete', data: any): void
    (e: 'onSaveUpdate', data: any): void
    (e: 'onSaveCreate', data: any): void
    (e: 'onSave', data: any): void
}