import { CommonModule } from '@angular/common';
import { Component, Input, NgModule } from '@angular/core';
import { FeedbackDTO } from '@pfeedback/feedback-data-spec';
import { UIBadgeModule } from './ui-badge.component';
import { UIVoteModule } from './ui-vote.component';

@Component({
  selector: 'app-ui-feedback-list',
  template: `
    <div class="py-6">
      <div
        *ngFor="let feedback of feedbackList"
        class="py-7 px-8 mb-5 flex space-x-10 bg-white rounded-lg"
      >
        <div class="flex-none">
          <app-ui-vote [votes]="feedback.votes"></app-ui-vote>
        </div>
        <div class="grow">
          <span class="mb-1 text-indigo-900 text-lg font-bold tracking-tight">{{
            feedback.title
          }}</span>
          <p class="mb-3 text-slate-600 text-base">
            {{ feedback.description }}
          </p>
          <app-ui-badge>{{ feedback.categoryName }}</app-ui-badge>
        </div>
      </div>
    </div>
  `,
})
export class UIFeedbackListComponent {
  @Input()
  feedbackList: FeedbackDTO[] = [];
}

@NgModule({
  imports: [CommonModule, UIVoteModule, UIBadgeModule],
  declarations: [UIFeedbackListComponent],
  exports: [UIFeedbackListComponent],
})
export class UIFeedbackListModule {}
