import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import {
  BoardDetailDTO,
  PFeedbackBoardSpecService,
} from '@pfeedback/board-data-spec';

@Injectable()
export class BoardFeedbacksHeadingService {
  constructor(private boardSpecService: PFeedbackBoardSpecService) {}

  getBoardHeadingDetails(boardId: string): Observable<BoardDetailDTO> {
    return this.boardSpecService.getBoardDetails(boardId);
  }
}
