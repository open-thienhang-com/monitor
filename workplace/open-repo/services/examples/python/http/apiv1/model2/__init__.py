from pmodules.adapters.http import TestNamespace, TestResource, TestHttpRespone

ns_model = TestNamespace(
    name="model2",
    description="Test V2"
)


@ns_model.route("")
@ns_model.doc(
    responses={
        200: "Success",
        400: "Bad request",
    },
)
class NamespaceTestV2(TestResource):
    def post(self):
        """Test namespace"""
        try:
            return 200
        except Exception as err:
            return 400
