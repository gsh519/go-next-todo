export type TodoType = {
    todo_id: number;
    content: string;
    is_done: boolean;
    deleted_at: string | null;
};

export class Todo {
    todoId: number;
    content: string;
    isDone: boolean;
    deletedAt: Date | null;

    constructor(properties?: TodoType) {
        this.todoId = properties ? properties.todo_id : 0;
        this.content = properties ? properties.content : '';
        this.isDone = properties ? properties.is_done : false;
        this.deletedAt = properties ? new Date(properties.deleted_at ?? "") : null;
    }
}
