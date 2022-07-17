import { CommonModule } from '@angular/common';
import {
  Component,
  EventEmitter,
  Input,
  NgModule,
  OnInit,
  Output,
} from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { FeedbackPayload } from '@pfeedback/feedback-data-spec';
import { UIFormFeaturePayload, DropdownOption } from './types';
import { UIDropdownModule } from './ui-dropdown.component';

@Component({
  selector: 'app-ui-create-feedback',
  template: `
    <form [formGroup]="form" class="relative">
      <div
        class="border rounded-full w-14 h-14 flex justify-center items-center bg-[url('assets/background-header.png')] bg-no-repeat bg-cover bg-center absolute -top-6 left-10"
      >
        <svg width="9" height="9" xmlns="http://www.w3.org/2000/svg">
          <text
            transform="translate(-24 -20)"
            fill="#F2F4FE"
            fill-rule="evenodd"
            font-family="Jost-Bold, Jost"
            font-size="14"
            font-weight="bold"
            class="fill-white"
          >
            <tspan x="24" y="27.5">+</tspan>
          </text>
        </svg>
      </div>
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
            formControlName="title"
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
            [options]="categoryOptions"
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
            formControlName="description"
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
            (click)="onSubmit()"
          >
            Add Feedback
          </button>
        </div>
      </div>
    </form>
  `,
})
export class UICreateFeedbackComponent implements OnInit {
  @Input()
  categoryOptions: DropdownOption[] = [];

  @Output()
  valuesToSubmit: EventEmitter<FeedbackPayload> = new EventEmitter<FeedbackPayload>();

  selectedCategory: number | null = null;
  form!: FormGroup;

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.form = this.fb.group<UIFormFeaturePayload>({
      title: '',
      description: '',
    });
  }

  onSelectCategory(value: string | number): void {
    this.selectedCategory = value as number;
  }

  onSubmit(): void {
    const formValues: FeedbackPayload = {
      ...this.form.value,
      categoryId: this.selectedCategory,
    };
    this.valuesToSubmit.emit(formValues);
  }
}

@NgModule({
  imports: [CommonModule, UIDropdownModule, ReactiveFormsModule],
  declarations: [UICreateFeedbackComponent],
  exports: [UICreateFeedbackComponent],
})
export class UICreateFeedbackModule {}
