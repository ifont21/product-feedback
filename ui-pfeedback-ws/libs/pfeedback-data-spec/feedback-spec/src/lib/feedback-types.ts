export interface FeedbackDTO {
  id: string;
  title: string;
  description?: string;
  votes: number;
  categoryName: string;
}

export interface FeedbackPayload {
  title: string;
  description?: string;
  categoryId: string;
}
