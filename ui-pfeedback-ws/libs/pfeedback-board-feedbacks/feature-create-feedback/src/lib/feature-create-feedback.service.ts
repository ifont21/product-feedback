import { Observable } from 'rxjs';
import {
  CategoryDTO,
  FeedbackPayload,
  PFeedbackBoardSpecService,
} from '@pfeedback/feedback-data-spec';
import { Injectable } from '@angular/core';

@Injectable()
export class FeatureCreateFeedbackService {
  constructor(private apiService: PFeedbackBoardSpecService) {}

  createNewFeedback(
    boardId: string,
    payload: FeedbackPayload
  ): Observable<unknown> {
    return this.apiService.createFeedback(boardId, payload);
  }

  getAllCategories(): Observable<CategoryDTO[]> {
    return this.apiService.getAllCategories();
  }
}
