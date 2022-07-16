import { FeedbackPayload } from '@pfeedback/feedback-data-spec';

export type UIFormFeaturePayload = Omit<FeedbackPayload, 'categoryId'>;

export interface DropdownOption {
  key: string | number;
  label: string;
}
