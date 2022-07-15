import { CommonModule } from '@angular/common';
import { Component, NgModule } from '@angular/core';

@Component({
  selector: 'app-ui-create-feedback',
  template: `
    <div>
      <h1>Create New Feedback</h1>
      <div>
        <label>Feedback Title</label>
        <p>Add a short, descriptive headline</p>
        <!-- input here -->
      </div>
      <div>
        <label>Category</label>
        <p>Choose a category for your feedback</p>
        <!-- Select drop down here -->
      </div>
      <div>
        <label>Feedback Detail</label>
        <p>
          Include any specific comments on what should be improved, added, etc.
        </p>
        <!-- Text area here -->
      </div>
      <div>
        <button>Cancel</button>
        <button>Add Feedback</button>
      </div>
    </div>
  `,
})
export class UICreateFeedbackComponent {}

@NgModule({
  imports: [CommonModule],
  declarations: [UICreateFeedbackComponent],
  exports: [UICreateFeedbackComponent],
})
export class UICreateFeedbackModule {}
