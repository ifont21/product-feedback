import { RouterModule } from '@angular/router';
import { UIFeedbackListModule } from './ui-feedback-list.component';
import { UIFeedbackListHeaderModule } from './ui-feedback-list-header.component';
import { FeatureFeedbackListService } from './feature-feedback-list.service';
import { Observable } from 'rxjs';
import { CommonModule } from '@angular/common';
import { Component, NgModule, OnInit } from '@angular/core';
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
  // TODO: should be fetch using the url
  public readonly boardId = 'b7c362c2-a617-4bf0-bf1f-014e0aec7bcb';

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
  imports: [
    CommonModule,
    UIFeedbackListHeaderModule,
    UIFeedbackListModule,
    RouterModule.forChild([
      {
        path: '',
        component: FeatureFeedbackListComponent,
      },
    ]),
  ],
  declarations: [FeatureFeedbackListComponent],
  exports: [FeatureFeedbackListComponent],
  providers: [FeatureFeedbackListService],
})
export class FeatureFeedbackListModule {}
