import { Component, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { BoardFeedbackHeadingContainerModule } from './board-feedbacks-heading-container.component';
import { BoardFeedbacksHeadingService } from './board-feedbacks-heading.service';

@Component({
  selector: 'app-board-feedback-shell',
  template: `<app-board-feedback-heading-container
    [boardId]="boardId"
  ></app-board-feedback-heading-container>`,
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
  ],
  declarations: [FeatureShellComponent],
  providers: [BoardFeedbacksHeadingService],
})
export class FeatureShellModule {}
