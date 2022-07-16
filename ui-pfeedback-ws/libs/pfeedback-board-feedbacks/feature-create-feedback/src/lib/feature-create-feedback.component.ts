import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { Component, NgModule, OnDestroy } from '@angular/core';
import { UICreateFeedbackModule } from './ui-create-feedback.component';
import { FeatureCreateFeedbackService } from './feature-create-feedback.service';
import { FeedbackPayload } from '@pfeedback/feedback-data-spec';
import { Subscription, Observable, map } from 'rxjs';
import { DropdownOption } from './types';

@Component({
  selector: 'app-create-feedback',
  template: `
    <div class="flex">
      <div>
        <a routerLink="../">Go Back</a>
        <div>
          <app-ui-create-feedback
            [categoryOptions]="(categories | async)!"
            (valuesToSubmit)="onSubmit($event)"
          ></app-ui-create-feedback>
        </div>
      </div>
    </div>
  `,
})
export class FeatureCreateFeedbackComponent implements OnDestroy {
  // TODO: should be fetch using the url
  public readonly boardId = 'b7c362c2-a617-4bf0-bf1f-014e0aec7bcb';

  subscription!: Subscription;

  categories: Observable<DropdownOption[]> = this.service
    .getAllCategories()
    .pipe(
      map((categories) =>
        categories.map(({ id, name }) => ({ key: id, label: name }))
      )
    );

  constructor(
    private service: FeatureCreateFeedbackService,
    private router: Router,
    private route: ActivatedRoute
  ) {}

  ngOnDestroy(): void {
    this.subscription?.unsubscribe();
  }

  onSubmit(values: FeedbackPayload): void {
    const submitSubscription = this.service
      .createNewFeedback(this.boardId, values)
      .subscribe({
        next: (success) => {
          this.router.navigate(['../'], { relativeTo: this.route });
        },
        error: (error) => {
          console.log(error);
        },
      });
    this.subscription.add(submitSubscription);
  }
}

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
  providers: [FeatureCreateFeedbackService],
  declarations: [FeatureCreateFeedbackComponent],
})
export class FeatureCreateFeedbackModule {}
