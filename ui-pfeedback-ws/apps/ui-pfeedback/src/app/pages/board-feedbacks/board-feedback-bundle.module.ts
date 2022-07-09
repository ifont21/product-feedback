import { RouterModule } from '@angular/router';
import { Component, NgModule } from '@angular/core';

@Component({
  selector: 'app-board-feedback-bundle',
  template: ` <div>Board Feedback Page!</div> `,
})
export class BoardFeedbackBundleComponent {}

const routes = [
  {
    path: '',
    component: BoardFeedbackBundleComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  declarations: [BoardFeedbackBundleComponent],
})
export class BoardFeedbackBundleModule {}
