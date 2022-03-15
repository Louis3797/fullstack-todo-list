export interface Todo {
  id: string;
  text: string;
  status: boolean;
  until?: number;
  created_at: number;
}
