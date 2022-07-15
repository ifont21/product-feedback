import {
  FeedbackDTO,
  PFeedbackBoardSpecService,
} from '@pfeedback/feedback-data-spec';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';

@Injectable()
export class FeatureFeedbackListService {
  constructor(private apiService: PFeedbackBoardSpecService) {}

  getBoardFeedbacks(boardId: string): Observable<FeedbackDTO[]> {
    return this.apiService.getBoardFeedbacks({ boardId });
  }
}
