import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { Component, NgModule } from '@angular/core';
import { UICreateFeedbackModule } from './ui-create-feedback.component';

@Component({
  selector: 'app-create-feedback',
  template: `
    <a routerLink="../">Go Back</a>
    <app-ui-create-feedback></app-ui-create-feedback>
  `,
})
export class FeatureCreateFeedbackComponent {}

@NgModule({
  imports: [
    CommonModule,
    UICreateFeedbackModule,
    RouterModule.forChild([
      {
        path: '',
        component: FeatureCreateFeedbackComponent,
      },
    ]),
  ],
  declarations: [FeatureCreateFeedbackComponent],
})
export class FeatureCreateFeedbackModule {}
