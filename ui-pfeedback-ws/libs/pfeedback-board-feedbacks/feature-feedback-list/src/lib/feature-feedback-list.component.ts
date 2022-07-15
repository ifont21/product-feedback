import { UIFeedbackListModule } from './ui-feedback-list.component';
import { UIFeedbackListHeaderModule } from './ui-feedback-list-header.component';
import { FeatureFeedbackListService } from './feature-feedback-list.service';
import { Observable } from 'rxjs';
import { CommonModule } from '@angular/common';
import { Component, Input, NgModule, OnInit } from '@angular/core';
import { FeedbackDTO } from '@pfeedback/feedback-data-spec';

@Component({
  selector: 'app-feature-feedback-list',
  template: `<div>
    <ng-container *ngIf="boardFeedbacks$ | async as boardFeedbacks">
      <app-ui-feedback-list-header
        [listCount]="boardFeedbacks?.length || 0"
      ></app-ui-feedback-list-header>
      <app-ui-feedback-list
        [feedbackList]="boardFeedbacks"
      ></app-ui-feedback-list>
    </ng-container>
  </div>`,
})
export class FeatureFeedbackListComponent implements OnInit {
  @Input()
  boardId = '';

  boardFeedbacks$: Observable<FeedbackDTO[]> | undefined;

  constructor(private feedbackListService: FeatureFeedbackListService) {}

  ngOnInit(): void {
    if (!this.boardId) return;

    this.boardFeedbacks$ = this.feedbackListService.getBoardFeedbacks(
      this.boardId
    );
  }
}

@NgModule({
  imports: [CommonModule, UIFeedbackListHeaderModule, UIFeedbackListModule],
  declarations: [FeatureFeedbackListComponent],
  exports: [FeatureFeedbackListComponent],
  providers: [FeatureFeedbackListService],
})
export class FeatureFeedbackListModule {}
