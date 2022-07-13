import { Component, Input, NgModule, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { BoardDetailDTO } from '@pfeedback/board-data-spec';
import { CommonModule } from '@angular/common';
import { BoardFeedbackHeadingModule } from './board-feedbacks-heading.component';
import { BoardFeedbacksHeadingService } from './board-feedbacks-heading.service';

@Component({
  selector: 'app-board-feedback-heading-container',
  template: `
    <app-board-feedbacks-heading
      [boardHeading]="boardHeading$ | async"
    ></app-board-feedbacks-heading>
  `,
})
export class BoardFeedbackHeadingContainerComponent implements OnInit {
  @Input()
  boardId: string | undefined;

  boardHeading$: Observable<BoardDetailDTO> | undefined;

  constructor(private service: BoardFeedbacksHeadingService) {}

  ngOnInit(): void {
    this.boardHeading$ = this.service.getBoardHeadingDetails(
      'b7c362c2-a617-4bf0-bf1f-014e0aec7bcb'
    );
  }
}

@NgModule({
  imports: [CommonModule, BoardFeedbackHeadingModule],
  declarations: [BoardFeedbackHeadingContainerComponent],
  exports: [BoardFeedbackHeadingContainerComponent],
})
export class BoardFeedbackHeadingContainerModule {}
