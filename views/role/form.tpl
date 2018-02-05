<form id="formIndex" method="post" role="form">
<div class="box">
    <div class="box-header">
        <div class="pull-left">
            <h3 class="box-title">Form Role</h3>
            <p>
                <br>
                Here you can set data user
            </p>
        </div>
    </div>
    <!-- /.box-header -->
    <div class="box-body">
        <div class="row">
            <div class="col-md-12">

                    <input value="" type="hidden" name="Id" />
                    <div class="form-group">
                        <label for="name">Name</label>
                        <input type="text" class="form-control" id="Name" name="Name" placeholder="Enter name" data-parsley-required="true">
                    </div>
                    <div class="form-group">
                        <label for="notes">Notes</label>
                        <textarea class="form-control" id="Notes" name="Notes" placeholder="Enter Notes"></textarea>
                    </div>
            </div>
        </div>
    </div>
    <div class="box-footer">
        <button type="button" class="btn bg-purple" onclick="submitFormIndex()"><i class="uk-icon-save"></i> Save</button>
        <button type="button" class="btn bg-yellow" data-dismiss="modal" aria-label="Close"><i class="uk-icon-arrow-circle-o-left"></i>  Close</button>
    </div>
</div>
</form>