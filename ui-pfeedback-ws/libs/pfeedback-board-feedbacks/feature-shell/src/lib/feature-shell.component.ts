import { FeatureFeedbackListModule } from '@pfeedback/board-feedbacks/feedback-list';
import { Component, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { BoardFeedbackHeadingContainerModule } from './board-feedbacks-heading-container.component';
import { BoardFeedbacksHeadingService } from './board-feedbacks-heading.service';

@Component({
  selector: 'app-board-feedback-shell',
  template: `
    <i class="bi-alarm"></i>
    <div class="flex space-x-7">
      <div class="flex-none w-64">
        <app-board-feedback-heading-container
          [boardId]="boardId"
        ></app-board-feedback-heading-container>
      </div>
      <div class="flex-auto">
        <app-feature-feedback-list
          [boardId]="boardId"
        ></app-feature-feedback-list>
      </div>
    </div>
  `,
})
export class FeatureShellComponent {
  // TODO: should be fetch using the url
  public readonly boardId = 'b7c362c2-a617-4bf0-bf1f-014e0aec7bcb';
}

@NgModule({
  imports: [
    RouterModule.forChild([
      {
        path: '',
        component: FeatureShellComponent,
      },
    ]),
    BoardFeedbackHeadingContainerModule,
    FeatureFeedbackListModule,
  ],
  declarations: [FeatureShellComponent],
  providers: [BoardFeedbacksHeadingService],
})
export class FeatureShellModule {}
