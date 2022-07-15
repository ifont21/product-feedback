import { Component, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { BoardFeedbackHeadingContainerModule } from './board-feedbacks-heading-container.component';
import { BoardFeedbacksHeadingService } from './board-feedbacks-heading.service';

@Component({
  selector: 'app-board-feedback-shell',
  template: `
    <div class="flex space-x-7">
      <div class="flex-none w-64">
        <app-board-feedback-heading-container
          [boardId]="boardId"
        ></app-board-feedback-heading-container>
      </div>

      <div class="flex-auto">
        <router-outlet></router-outlet>
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
        children: [
          {
            path: '',
            loadChildren: () =>
              import('@pfeedback/board-feedbacks/feedback-list').then(
                (m) => m.FeatureFeedbackListModule
              ),
          },
          {
            path: 'create',
            loadChildren: () =>
              import('@pfeedback/board-feedback/create-feedback').then(
                (m) => m.FeatureCreateFeedbackModule
              ),
          },
        ],
      },
    ]),
    BoardFeedbackHeadingContainerModule,
  ],
  declarations: [FeatureShellComponent],
  providers: [BoardFeedbacksHeadingService],
})
export class FeatureShellModule {}
