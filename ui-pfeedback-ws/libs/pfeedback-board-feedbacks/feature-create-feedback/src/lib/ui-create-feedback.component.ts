import { CommonModule } from '@angular/common';
import { Component, NgModule } from '@angular/core';
import { UIDropdownModule } from './ui-dropdown.component';

@Component({
  selector: 'app-ui-create-feedback',
  template: `
    <div class="rounded-lg p-10 bg-white">
      <h1 class="text-2xl font-bold tracking-tight text-indigo-900 mb-10">
        Create New Feedback
      </h1>
      <div class="mb-6">
        <label for="title">
          <span class="text-sm font-bold tracking-tight text-indigo-900"
            >Feedback Title</span
          >
          <p class="text-slate-600 text-base mb-4">
            Add a short, descriptive headline
          </p>
        </label>
        <input
          type="text"
          id="title"
          class="h-12 w-full rounded bg-slate-100"
        />
      </div>

      <div class="mb-6">
        <label for="category">
          <span class="text-sm font-bold tracking-tight text-indigo-900"
            >Category</span
          >
          <p class="text-slate-600 text-base mb-4">
            Choose a category for your feedback
          </p>
        </label>
        <app-ui-dropdown
          [options]="options"
          (selectedItem)="onSelectCategory($event)"
        ></app-ui-dropdown>
      </div>

      <div class="mb-8">
        <label for="description">
          <span class="text-sm font-bold tracking-tight text-indigo-900"
            >Feedback Detail</span
          >
          <p class="text-slate-600 text-base mb-4">
            Include any specific comments on what should be improved, added,
            etc.
          </p>
        </label>
        <!-- Text area here -->
        <textarea
          id="description"
          cols="30"
          rows="10"
          class="w-full rounded resize-none bg-slate-100"
        ></textarea>
      </div>
      <div class="flex justify-end">
        <button
          class="rounded-lg px-6 py-3 bg-indigo-900 text-white font-bold text-sm mr-4"
        >
          Cancel
        </button>
        <button
          class="rounded-lg px-6 py-3 bg-fuchsia-700 text-white font-bold text-sm"
        >
          Add Feedback
        </button>
      </div>
    </div>
  `,
})
export class UICreateFeedbackComponent {
  options = [
    { key: 'Option 1', label: 'Option 1' },
    { key: 'Option 2', label: 'Option 2' },
    { key: 'Option 3', label: 'Option 3' },
  ];

  selectedCategory = '';

  onSelectCategory(value: string | number): void {
    this.selectedCategory = value as string;
  }
}

@NgModule({
  imports: [CommonModule, UIDropdownModule],
  declarations: [UICreateFeedbackComponent],
  exports: [UICreateFeedbackComponent],
})
export class UICreateFeedbackModule {}
