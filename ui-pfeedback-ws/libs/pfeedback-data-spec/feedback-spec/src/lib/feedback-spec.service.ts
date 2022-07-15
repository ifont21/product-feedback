import { Inject, Injectable, Optional } from '@angular/core';
import { SpecConfig, Configuration } from './config';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { FeedbackDTO } from './feedback-types';

@Injectable({
  providedIn: 'root',
})
export class PFeedbackBoardSpecService {
  constructor(
    @Optional() @Inject(SpecConfig) private config: Configuration,
    private http: HttpClient
  ) {}

  getBoardFeedbacks({
    boardId,
  }: {
    boardId: string;
  }): Observable<FeedbackDTO[]> {
    return this.http.get<FeedbackDTO[]>(
      `${this.config.basePath}/boards/${boardId}/feedbacks`
    );
  }
}
