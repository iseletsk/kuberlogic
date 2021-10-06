import { NO_ERRORS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ChangePasswordDialogComponent } from '@pages/services-page/pages/view-service/pages/service-connection/components/change-password-dialog/change-password-dialog.component';
import { FormBuilder } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MockMatDialogRef } from '@testing/mock-mat-dialog-ref';


describe('ChangePasswordDialogComponent', () => {
    let component: ChangePasswordDialogComponent;
    let fixture: ComponentFixture<ChangePasswordDialogComponent>;
    let dialogRef: MockMatDialogRef;
    function setFormData(data: { password: string, confirmPassword: string}) {
        component.formGroup.patchValue(data);
        fixture.detectChanges();
    }
    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ChangePasswordDialogComponent],
            providers: [
                FormBuilder,
                { provide: MatDialogRef, useClass: MockMatDialogRef },
            ],
            schemas: [NO_ERRORS_SCHEMA]
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ChangePasswordDialogComponent);
        component = fixture.componentInstance;
        // @ts-ignore
        dialogRef = TestBed.inject(MatDialogRef);
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    describe('when form is invalid', () => {
        it('should not emit form value when passwords must match', () => {
            setFormData({password: 'testPass1', confirmPassword: 'testPass2'});
            const spy = spyOn(dialogRef, 'close');
            component.onSave();
            fixture.detectChanges();

            expect(spy).not.toHaveBeenCalled();
        });

        it('should not emit form value when fields are empty', () => {
            setFormData({password: '', confirmPassword: ''});
            const spy = spyOn(dialogRef, 'close');
            component.onSave();
            fixture.detectChanges();

            expect(spy).not.toHaveBeenCalled();
        });
    });

    describe('when form is valid', () => {
        it('should emit password field value', () => {
            setFormData({password: 'testPass1', confirmPassword: 'testPass1'});
            const spy = spyOn(dialogRef, 'close');
            component.onSave();
            fixture.detectChanges();

            expect(spy).toHaveBeenCalledWith('testPass1');
        });
    });
});
